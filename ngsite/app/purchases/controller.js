'use strict';

angular.module('bestbuys.purchases', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/purchases', {
    templateUrl: 'purchases/template.html',
    controller: 'PurchasesCtrl'
  });
}])

.controller('PurchasesCtrl', ['$scope', function($scope) {
  $scope.gameId = '54e2ac43f047051f3c000004';    
  
  $scope.generatePurchases = function() {
  };
}]); 


	  