'use strict';

var gameServices = angular.module('bestbuys.services', ['ngResource']);

gameServices.factory('Game', ['$resource',
	 function($resource){
		 return $resource('/games/:id', {id:'@id'});
	 }]);
