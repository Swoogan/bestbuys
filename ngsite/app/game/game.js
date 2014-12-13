'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', '$http', function($scope, $http) {
  //$http.get('/games/54177f9ef047050f92000004').
  $http.get('/games/54177f9ef047050f9200000').then(
    function(result) {
      $scope.finance = result.data;
    },
    function(error) {
      console.log(error.statusText);
      if (error.status == 404)
	$scope.error = "The requested game data could not be found";
      else
	$scope.error = error.statusText;
      /*
      $(this).removeClass('success');
      $(this).addClass('error');
      $(this).fadeToggle('slow');
      $('#error').text("Error in: '" + settings.url + "'");
      $('#exception').text('Message: ' + exception);	  
      */
    });
}]); 
