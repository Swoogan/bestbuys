'use strict';

describe('bestbuys.purchases module', function() {

  beforeEach(module('bestbuys.purchases'));

  describe('purchases controller', function(){
    var scope, ctrl;
    
    beforeEach(inject(function($rootScope, $controller) {      
      scope = $rootScope.$new();
      ctrl = $controller('PurchasesCtrl', {$scope: scope});
    }));

    it('should ....', function() {
      expect(ctrl).toBeDefined();
    });
    
    /*
    it('should create "structures" model with 2 structure', function() {
      expect(scope.finance.structures.length).toBe(11);
    });
    
    it('should have finance', function() {
      expect(scope.finance).toBeDefined();
    });
    
    it('should have income 0', function() {
      expect(scope.finance.income).toBe(0);
    });   
    */
  });
}); 
