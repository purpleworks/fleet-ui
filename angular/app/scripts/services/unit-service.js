'use strict';

/**
 * @ngdoc service
 * @name fleetuiApp.unitService
 * @description
 * # unitService
 * Service in the fleetuiApp.
 */
angular.module('fleetuiApp')
  .service('unitService', function unitService($resource) {
    return $resource('/api/v1/units/:id', {
      id: '@id'
    });
  });
