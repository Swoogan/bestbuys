'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', function($scope) {
   $scope.structures = [
    {
      'name': 'Supply Depot',
      'cost': '2000000',
      'increase': '15000',
      'income': '3000'
    },
    {
      'name': 'Collection Outpost',
      'cost': '5000000',
      'increase': '75000',
      'income': '10000'
    }
  ];  
  
  $scope.finance = {
    'income': 0,
    'upkeep': 0,
    'balance': 0,
    'wallet': 0,
    'lands': 0
    
  };
  
  $scope.class = '';
  $scope.editFocus = function() {    
    $scope.class = 'shadow';
  }
}]); 
