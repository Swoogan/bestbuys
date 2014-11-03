'use strict';

angular.module('bestbuys.version', [
  'bestbuys.version.interpolate-filter',
  'bestbuys.version.version-directive'
])

.value('version', '0.1');
