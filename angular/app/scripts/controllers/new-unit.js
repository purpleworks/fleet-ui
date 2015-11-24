'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:NewUnitCtrl
 * @description
 * # NewUnitCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('NewUnitCtrl', function ($scope, $state, $timeout, unitService) {
    $scope.newUnit = {
      isLoading: false,
      name: 'new-service.service',
      service: "[Unit]\n" +
        "Description=New Service\n" +
        "Requires=docker.service\n" +
        "After=docker.service\n" +
        "\n[Service]\n" +
        "ExecStart=\n" +
        "ExecStop=\n" +
        "\n[X-Fleet]\n",
      submit: function() {
        var _this = this;
        if(!_this.isLoading && confirm('Start unit?')) {
          _this.isLoading = true;
          unitService.save({ name: _this.name, service: _this.service }, function(data) {
            _this.isLoading = false;
            $state.go('root.main.unit', { name: _this.name });
          }, function() {
            _this.isLoading = false;
          });
        }
      }
    }

    $scope.unitUpload = {
      loading: false,
      url: 'api/v1/units/upload',
      options: {
        container: 'unitUploadBtn',
        multi_selection: false,
        max_file_size: '6mb',
        filters: {
          mime_types : [
            { title: "service file", extensions : "service" }
          ]
        }
      },
      callbacks: {
        filesAdded: function(uploader, files) {
          $scope.unitUpload.loading = true;
          $timeout(function() { uploader.start() }, 1);
        },
        fileUploaded: function(uploader, file, response) {
          $scope.unitUpload.loading = false;

          if(response.response) {
            var data = $.parseJSON(response.response);
            if(data.result == 'success') {
              $state.go('root.main.unit', { name: file.name });
            }
          }
        },
        error: function(uploader, error) {
          $scope.unitUpload.loading = false;

          if(error.code == plupload.FILE_SIZE_ERROR) {
            alert("최대 6mb까지 첨부가능합니다.");
          } else {
            alert(error.message);
          }
        }
      }
    };
  });
