'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:UnitCtrl
 * @description
 * # UnitCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('UnitCtrl', function ($scope, $state, unitService) {
    $scope.unit = unitService.get({ id:$state.params.name });
  });
