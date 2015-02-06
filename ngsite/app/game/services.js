'use strict';

var gameServices = angular.module('bestbuys.services', ['ngResource']);

gameServices.factory('Game', ['$resource',
	 function($resource){
		 return $resource('games/:id', {}, {
			 query: {method:'GET', params:{id:'54d42204f047050fc600000a'}}
		 });
	 }]);
