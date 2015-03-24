'use strict';

/*global angular*/

var bestbuys = angular.module('bestbuys');

bestbuys.directive('bbCost', function ($filter, utilities) {
  return {
    restrict: 'A',
    replace: true,
    scope: { cost: '=', structure: '@', save: '=' },
    template: '<td class="money editable" contentEditable="true">{{cost | currency:undefined:0}}</td>',
    link: function (scope, element) {
      /* On Focus */
      element.bind('focus', function () {
        element.addClass('shadow');
        element.html(utilities.parseCurrency(element.html()));
      });

      /* On Blur */
      element.bind('blur', function () {
        var amount = element.html();
        element.removeClass('shadow');
        element.html($filter('currency')(amount, undefined, 0));
        scope.save(scope.structure, utilities.parseCurrency(amount));
      });
 
      /* On Keydown */
      element.bind('keydown', utilities.handleEnter);

      /* On Keyup */
      element.bind('keyup', utilities.handleEscape);
    }
  };
});
