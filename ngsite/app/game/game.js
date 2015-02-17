'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', '$routeParams', 'Game', 'Command', 'ErrorMessage',
  function($scope, $routeParams, Game, Command, ErrorMessage) {
    $scope.gameId = '54d42204f047050fc600000a';
    
    Game.get({id: $scope.gameId},
      function(data) {
	$scope.finance = data;
      },
      function(error) {
	$scope.message = ErrorMessage.generate(error);      
      }
    );
    
    $scope.saveValue = function(field, value) {    
      var command = 'set' + field.charAt(0).toUpperCase() + field.slice(1);
      var data = { game: $scope.gameId };
      data[field] = value;
      
      Command.save(command, data);
    }
    
    $scope.saveStructure = function(name, value) {        
      var data = { structureCost: value, structureName: name, game: $scope.gameId };
      Command.save('setStructureCost', data);
    }    
  }
]); 


