'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('MainCtrl', function ($scope, $interval, machineService, unitService) {
    $scope.machine = {
      loading: false,
      items: machineService.query()
    };

    $scope.unit = {
      loading: false,
      items: {}
    };

    function getUnitsInfo(init) {
      unitService.query(function(data) {
        $scope.unit.items = {};
        for(var i=0; i<data.length; i++) {
          var unit = data[i];
          var machineName = unit.Machine.split("/")[0];
          if($scope.unit.items[machineName]) {
            $scope.unit.items[machineName].push(unit);
          } else {
            $scope.unit.items[machineName] = [unit];
          }
        }
      });
    }

    $interval(getUnitsInfo, 5000);
    getUnitsInfo();
  });
