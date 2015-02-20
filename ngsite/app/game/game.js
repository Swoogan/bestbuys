'use strict';

angular.module('bestbuys.game', ['ngRoute', 'ngResource'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', 'Game', 'Command', 'Notification',
  function($scope, Game, Command, Notification) {
    $scope.gameId = '54e6a1a7f047055eb3000004';    
    
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
]) 

.factory('Game', ['$resource', function($resource) {
    return $resource('/games/:id', {id:'@id'});
}])

.service('Notification', [function() {  
  this.showSuccess = function () {	      
    return {
      error: false,
      show: true,
      text: 'Successfully saved changes.'
    };      
  };
  
  this.showError = function(error) {
    var message = {
      error: true,
      show: true,
      text: 'Error in: ' + error.config.url
    };
    
    if (error.status == 404)
      message.details = "The requested game data could not be found.";
    else if (error.status == 502)
      message.details = "The game data service appears unavailable. Try again later.";
    else
      message.details = error.statusText;
    
    return message;
  };
}])
  
.service('Command', ['$http', function($http) {   
  this.save = function (name, data) {
    return $http.post('/commands/', {name: name, data: data});
  };  
}]);
