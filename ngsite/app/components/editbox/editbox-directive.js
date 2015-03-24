'use strict';

/*global angular*/

var bestbuys = angular.module('bestbuys');

bestbuys.directive('bbEditbox', function ($filter, parseCurrency) {
  return {
    restrict: 'E',
    replace: true,
    scope: { value: '=', name: '@', save: '=' },
    templateUrl: '/app/components/editbox/editbox.html',
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

