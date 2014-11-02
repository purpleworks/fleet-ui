'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:NewUnitCtrl
 * @description
 * # NewUnitCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('NewUnitCtrl', function ($scope, $state, unitService) {
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
  });
