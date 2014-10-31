'use strict';

/**
 * @ngdoc overview
 * @name fleetuiApp
 * @description
 * # fleetuiApp
 *
 * Main module of the application.
 */
angular
  .module('fleetuiApp', [
    'config',
    'ngAnimate',
    'ngCookies',
    'ngResource',
    'ngSanitize',
    'ngTouch',
    'ui.router'
  ])
  .config(function ($stateProvider, $urlRouterProvider, CACHE_VERSION) {
    // prevent view cache helper
    function _t(url) {
      return url + '?_cache=' + CACHE_VERSION;
    }

    // default route
    $urlRouterProvider.when('', '/');

    // route
    $stateProvider
      /* login 하기 전 */
      .state('root', {
        abstract: true,
        templateUrl: _t('/views/layout.html')
      })
        .state('root.main', {
          url: '/',
          templateUrl: _t('/views/main.html'),
          controller: 'MainCtrl'
        });
  })
  .run(function($rootScope, $state, $stateParams) {
    // set default root scope value
    $rootScope.$state = $state;
    $rootScope.$stateParams = $stateParams;
  });
