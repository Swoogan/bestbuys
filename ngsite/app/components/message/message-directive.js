'use strict';

/*global angular*/

var bestbuys = angular.module('bestbuys');

bestbuys.directive('bbMessage', function () {
  return {
    link: function (scope, element) {
      element.bind('click', function () {
        scope.message.show = false;
        scope.$apply();
      });

      scope.$watch("message.error", function () {
        if (scope.message && scope.message.error) {
          element.removeClass('success');
          element.addClass('error');
        } else {
          element.removeClass('error');
          element.addClass('success');
        }
      });
    }
  };
});

