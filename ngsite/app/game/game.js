'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', '$http', function($scope, $http) {
  $scope.gameId = '548f1a5df047051df5000005';
  //$scope.gameId = '54177f9ef047050f9200000'; 
  
  $http.get('/games/' + $scope.gameId).then(    
    function(result) {
      $scope.finance = result.data;      
    },    
    function(error) {
      $scope.message = {};
      $scope.message.show = true;
      errorMessage($scope.message, error);
    }
  );
  
  $scope.saveValue = function(field, value) {    
    var command = 'set' + field.charAt(0).toUpperCase() + field.slice(1).toLowerCase()    
    var data = { game: $scope.gameId };
    data[field] = value;   
    
    $scope.message = {};
    $scope.message.show = true;
    
    $http.post('/commands/', {name: command, data: data}).then(
      function () {	
	$scope.message.error = false;	
	$scope.message.text = 'Successfully saved changes.';	
      },
      function(error) {	
	errorMessage($scope.message, error);
      }
    );
  }
  
  $scope.saveStructure = function(name, value) {        
    var data = { structureCost: value, structureName: name, game: $scope.gameId };
    $http.post('/commands/', {name: 'setStructureCost', data: data}).then(
      function () {
	$scope.message = {};
	$scope.message.error = false;
	$scope.message.show = true;
	$scope.message.text = 'Successfully saved changes.';	
      },
      function(error) {	
	errorMessage($scope.message, error);
      }
    );
  }
}]); 

function errorMessage(message, error) {
  message.error = true;
  message.text = 'Error in: ' + error.config.url;
  
  if (error.status == 404)
    message.details = "The requested game data could not be found.";
  else if (error.status == 502)
    message.details = "The game data service appears unavailable. Try again later.";
  else
    message.details = error.statusText;
}