'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:UnitCtrl
 * @description
 * # UnitCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('UnitCtrl', function ($scope, $state, $location, $timeout, unitService, WebSocket, ENVIRONMENT) {
    var unitName = $state.params.name;

    var timer = null;

    $scope.logs = [];

    function getUnitInfo() {
      unitService.get({ id: unitName }, function(data) {
        $scope.unit = data;

        if(timer) {
          $timeout.cancel(timer)
        }
        timer = $timeout(function() {
          getUnitInfo();
        }, 1000);
      });
    }

    getUnitInfo();
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

    $scope.destroyUnit = function() {
      $scope.destroyLoading = true;
      if(confirm('Destroy service?')) {
        unitService.remove({ id:$state.params.name }, function(data) {
          $scope.destroyLoading = false;
        }, function() {
          $scope.destroyLoading = false;
        });
      }
    }

    if(ENVIRONMENT == 'dev') {
      WebSocket.new('ws://' + $location.host() + ':3000' + '/ws/journal/' + unitName);
    } else {
      WebSocket.new('ws://' + $location.host() + ':' + $location.port() + '/ws/journal/' + unitName);
    }
    setCallback();
    $scope.$on('$destroy', function () {
      WebSocket.close();
    });
  });
