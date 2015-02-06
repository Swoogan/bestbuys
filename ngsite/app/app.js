'use strict';

// Declare app level module which depends on views, and components
var bestbuyApp = angular.module('bestbuys', [
  'ngRoute',
  'ngAnimate',
  'bestbuys.game',
  'bestbuys.version',
  'bestbuys.services'
]);

bestbuyApp.config(['$routeProvider', function($routeProvider) {
  $routeProvider.otherwise({redirectTo: '/game'});
}]);

bestbuyApp.directive('bbEditbox', function ($filter) {
  return {    
    restrict: 'E',
    replace: true,
    scope: { value: '=', name: '@', label: '@', save: '=' },
    templateUrl: 'bb-editbox.html',
    link: function (scope, element, attrs) {
      // the second div is always the actual editbox (see the template)
      var editbox = angular.element(element.find('div')[1]);
      
      /* On Focus */
      editbox.bind('focus', function () {
        editbox.addClass('shadow');
        editbox.html(parseCurrency(editbox.html()));
      });
      
      /* On Blur */
      editbox.bind('blur', function () {
        var amount = editbox.html();
        editbox.removeClass('shadow');
        editbox.html($filter('currency')(amount, undefined, 0));
        scope.save(attrs.name, parseCurrency(amount));
      });
      
      /* On Keydown */
      editbox.bind('keydown', handleEnter);
      
      /* On Keyup */
      editbox.bind('keyup', handleEscape);
    }    
  }
});

bestbuyApp.directive('bbCost', function ($filter) {
  return {
    restrict: 'A',
    replace: true,
    scope: { cost: '=', structure: '@', save: '=' },
    template: '<td class="money editable" contentEditable="true">{{cost | currency:undefined:0}}</td>',
    link: function (scope, element, attrs) {
      /* On Focus */
      element.bind('focus', function () {
        element.addClass('shadow');
        element.html(parseCurrency(element.html()));
      });
      
      /* On Blur */
      element.bind('blur', function () {
        var amount = element.html();
        element.removeClass('shadow');
        element.html($filter('currency')(amount, undefined, 0));        
        scope.save(scope.structure, parseCurrency(amount));
      });
      
      /* On Keydown */
      element.bind('keydown', handleEnter);
      
      /* On Keyup */
      element.bind('keyup', handleEscape);
    }    
  }
});

bestbuyApp.directive('bbMessage', function ($animate) {
  return {
    link: function (scope, element, attrs) {      
      element.bind('click', function () {        
        scope.message.show = false;
        scope.$apply();
      });

      scope.$watch("message.error", function() {
        if (scope.message && scope.message.error) {
          element.removeClass('success');
          element.addClass('error');
        }
        else {
          element.removeClass('error');
          element.addClass('success');        
        }
      });
    }    
  }
});

bestbuyApp.filter('capitalize', function() {
  return function(input) {
    return (!!input) ? input.charAt(0).toUpperCase() + input.slice(1).toLowerCase() : '';
  }
});


function parseCurrency(amount) {
  if (amount === "") return 0;
  
  if (amount.charAt(0) == "$")   
    amount = amount.slice(1);
  
  amount = amount.replace(/<br>/g, "");
  amount = amount.replace(/,/g, "");
  var i = parseInt(amount);
  
  if (isNaN(i)) return "0";
  
  return i;
}

function handleEscape(e) {
  var esc = e.which == 27;
  
  if (esc) {
    document.execCommand("undo", false, null);
    e.target.blur();
  } 
}

function handleEnter(e) {
  var enter = e.which == 13;

  if (enter) {
    e.preventDefault();
    e.target.blur();
  }
}
