'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/template.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', 'Game', 'Command', 'Notification',
  function($scope, Game, Command, Notification) {
    $scope.gameId = '54e2ac43f047051f3c000004';    
    
    Game.get({id: $scope.gameId},
      function(data) {
        $scope.finance = data;
      },
      function (error) {
        $scope.message = Notification.showError(error);
      }
    );
    
    $scope.saveValue = function(field, value) {    
      var command = 'set' + field.charAt(0).toUpperCase() + field.slice(1);
      var data = { game: $scope.gameId };
      data[field] = value;
      
      Command.save(command, data).then(
        function () {
          $scope.message = Notification.showSuccess();
        },
        function (error) {
          $scope.message = Notification.showError(error);
        }
      );
    }
    
    $scope.saveStructure = function(name, value) {        
      var data = { structureCost: value, structureName: name, game: $scope.gameId };
      Command.save('setStructureCost', data).then(
        function () {
          $scope.message = Notification.showSuccess();
        },
        function (error) {
          $scope.message = Notification.showError(error);
        }
      );
    }    
  }
]); 


	  