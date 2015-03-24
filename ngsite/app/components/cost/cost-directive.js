'use strict';

/*global angular*/

var bestbuys = angular.module('bestbuys');

bestbuys.directive('bbCost', function ($filter, parseCurrency) {
  return {
    restrict: 'A',
    replace: true,
    scope: { cost: '=', structure: '@', save: '=' },
    template: '<td class="money editable" contentEditable="true">{{cost | currency:undefined:0}}</td>',
    link: function (scope, element) {
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
  };
});
