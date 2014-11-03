'use strict';

describe('bestbuys.version module', function() {
  beforeEach(module('bestbuys.version'));

  describe('version service', function() {
    it('should return current version', inject(function(version) {
      expect(version).toEqual('0.1');
    }));
  });
});
