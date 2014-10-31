'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('MainCtrl', function ($scope, machineService, unitService) {
    $scope.machine = {
      loading: false,
      items: machineService.query()
    };

    $scope.unit = {
      loading: false,
      items: unitService.query()
    };
  });
