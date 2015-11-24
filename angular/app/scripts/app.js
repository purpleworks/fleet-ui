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
    'ui.router',
    'angular-websocket',
    'angular-ladda',
    'angular-plupload'
  ])
  .config(function ($stateProvider, $urlRouterProvider, pluploadOptionProvider, CACHE_VERSION) {
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
        templateUrl: _t('views/layout.html')
      })
        .state('root.main', {
          url: '/',
          templateUrl: _t('views/main.html'),
          controller: 'MainCtrl'
        })
          .state('root.main.unit', {
            url: 'unit?name',
            templateUrl: _t('views/unit.html'),
            controller: 'UnitCtrl'
          })
          .state('root.main.new_unit', {
            url: 'unit/new',
            templateUrl: _t('views/new-unit.html'),
            controller: 'NewUnitCtrl'
          });

    // Plupload 설정
    /* jshint camelcase:false */
    pluploadOptionProvider.setOptions({
      flash_swf_url: 'bower_components/plupload/js/Moxie.swf',
      silverlight_xap_url: 'bower_components/plupload/js/Moxie.xap'
    });
  })
  .run(function($rootScope, $state, $stateParams) {
    // set default root scope value
    $rootScope.$state = $state;
    $rootScope.$stateParams = $stateParams;
  });
