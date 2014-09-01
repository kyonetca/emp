'use strict';

/**
 * @ngdoc function
 * @name empAngularApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the empAngularApp
 */
angular.module('empAngularApp')
  .controller('MainCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
