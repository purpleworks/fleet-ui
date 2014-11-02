'use strict';

/**
 * @ngdoc filter
 * @name fleetuiApp.filter:serviceName
 * @function
 * @description
 * # serviceName
 * Filter in the fleetuiApp.
 */
angular.module('fleetuiApp')
  .filter('serviceName', function () {
    return function (input) {
      if(input) {
        return input.substring(0, input.lastIndexOf('.service'));
      } else {
        return null;
      }
    };
  });
