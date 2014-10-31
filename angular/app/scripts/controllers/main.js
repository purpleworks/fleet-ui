'use strict';

/**
 * @ngdoc function
 * @name fleetuiApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the fleetuiApp
 */
angular.module('fleetuiApp')
  .controller('MainCtrl', function ($scope, unitService) {
    $scope.units = unitService.query();
  });
