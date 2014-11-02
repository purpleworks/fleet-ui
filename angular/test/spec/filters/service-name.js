'use strict';

describe('Filter: serviceName', function () {

  // load the filter's module
  beforeEach(module('fleetuiApp'));

  // initialize a new instance of the filter before each test
  var serviceName;
  beforeEach(inject(function ($filter) {
    serviceName = $filter('serviceName');
  }));

  it('should return the input prefixed with "serviceName filter:"', function () {
    var text = 'angularjs';
    expect(serviceName(text)).toBe('serviceName filter: ' + text);
  });

});
