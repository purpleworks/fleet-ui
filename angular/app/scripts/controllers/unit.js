'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:UnitCtrl
 * @description
 * # UnitCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('UnitCtrl', function ($scope, $state, $location, $timeout, unitService, WebSocket) {
    var timer = null;

    $scope.logs = [];

    function getUnitInfo(init) {
      unitService.get({ id:$state.params.name }, function(data) {
        $scope.unit = data;

        if(init) {
          WebSocket.new('ws://' + $location.host() + ':3000' + '/ws/journal/' + data.Unit);
          setCallback();
          $scope.$on("$destroy", function () {
            WebSocket.close();
          });
        }

        if(timer) {
          $timeout.cancel(timer)
        }
        timer = $timeout(function() {
          getUnitInfo();
        }, 1000);
      });
    }

    getUnitInfo(true);

    $scope.$on("$destroy", function() {
      $timeout.cancel(timer);
    });

    function setCallback() {
      // recv
      WebSocket.onmessage(function (e) {
        $scope.logs.push(e.data);
        $timeout(function() {
          $('.unit-log').scrollTop($('.unit-log')[0].scrollHeight);
        }, 50);
      });

      // connection open
      WebSocket.onopen(function () {
        console.log('%cWebSocket Connection Open!', 'background: #34495e; color: #fff');
      });

      // connection close
      WebSocket.onclose(function () {
        console.log('%cWebSocket Connection Close!', 'background: #34495e; color: #fff');
      });

      // error
      WebSocket.onerror(function (e) {
        console.log('%cWebSocket Error Occur!', 'background: #34495e; color: #fff');
      });
    }

    $scope.startUnit = function() {
      $scope.startLoading = true;
      unitService.start({ id:$state.params.name }, function(data) {
        $scope.startLoading = false;
      }, function() {
        $scope.startLoading = false;
      });
    };

    $scope.stopUnit = function() {
      $scope.stopLoading = true;
      unitService.stop({ id:$state.params.name }, function(data) {
        $scope.stopLoading = false;
      }, function() {
        $scope.stopLoading = false;
      });
    };

    $scope.forceStopUnit = function() {
      $scope.forceStopLoading = true;
      unitService.stop({ id:$state.params.name }, function(data) {
        $scope.forceStopLoading = false;
      }, function() {
        $scope.forceStopLoading = false;
      });
    };
  });
