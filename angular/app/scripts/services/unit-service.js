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
    return $resource('api/v1/units/:collection_action/:id/:member_action', {
      id: '@id'
    }, {
      start: {
        method: 'POST',
        params: {
          member_action: 'start'
        }
      },
      stop: {
        method: 'POST',
        params: {
          member_action: 'stop'
        }
      }
    });
  });
