'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', '$http', function($scope, $http) {
  $scope.gameId = '54177f9ef047050f92000004';
  //$scope.gameId = '54177f9ef047050f9200000';
  
  $http.get('/games/' + $scope.gameId).then(
    function(result) {
      $scope.finance = result.data;      
    },    
    function(error) {
      $scope.message = errorMessage(error);      
    }
  );
  
  $scope.saveValue = function(field, value) {
    var http = $http;
    
    var command = 'set' + field.charAt(0).toUpperCase() + field.slice(1).toLowerCase()
    
    //if (focusoutEnabled) {       
      var data = { game: $scope.gameId };
      data[field] = value;
      post(http, command, data);
    //}
  }
  
  $scope.saveStructure = function(name, value) {    
    //if (focusoutEnabled) {       
      var data = '{"structureCost": ' + value + ', "structureName": "' + name + '", "game": "' + $scope.gameId +'"}';
      post(http, 'setStructureCost', data);      
    //}    
  }
}]); 

function post(http, command, data) {
  http.post('/commands/', {name: command, data: data}).then(
    function(error) {
      cosole.log(error);
      //$scope.message = errorMessage(error);
    }
  );
}

function errorMessage(error) {
  message = {};
      
  message.error = true;
  message.text = 'Error in: ' + error.config.url;
  
  if (error.status == 404)
    message.details = "The requested game data could not be found.";
  else if (error.status == 502)
    message.details = "The game data service appears unavailable. Try again later.";
  else
    message.details = error.statusText;
  
  message.show = true;
  
  return message;
}