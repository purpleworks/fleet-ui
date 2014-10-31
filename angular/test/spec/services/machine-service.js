'use strict';

describe('Service: machineService', function () {

  // load the service's module
  beforeEach(module('fleetuiApp'));

  // instantiate service
  var machineService;
  beforeEach(inject(function (_machineService_) {
    machineService = _machineService_;
  }));

  it('should do something', function () {
    expect(!!machineService).toBe(true);
  });

});
