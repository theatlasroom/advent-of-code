// Generated by BUCKLESCRIPT VERSION 4.0.6, PLEASE EDIT WITH CARE
'use strict';

var $$Array = require("bs-platform/lib/js/array.js");

function str_to_list($staropt$star, input) {
  var delimiter = $staropt$star !== undefined ? $staropt$star : "\n";
  return $$Array.to_list(input.trim().split(delimiter));
}

var Utils = /* module */[/* str_to_list */str_to_list];

exports.Utils = Utils;
/* No side effect */
