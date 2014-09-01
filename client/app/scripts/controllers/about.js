'use strict';

/**
 * @ngdoc function
 * @name empAngularApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the empAngularApp
 */
angular.module('empAngularApp')
  .controller('AboutCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
