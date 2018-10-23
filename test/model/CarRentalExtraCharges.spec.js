/**
 * Payment Gateway API Specification
 * Payment Gateway API for payment processing. 
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: unset
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.FirstDataGateway);
  }
}(this, function(expect, FirstDataGateway) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new FirstDataGateway.CarRentalExtraCharges();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('CarRentalExtraCharges', function() {
    it('should create an instance of CarRentalExtraCharges', function() {
      // uncomment below and update the code to test CarRentalExtraCharges
      //var instane = new FirstDataGateway.CarRentalExtraCharges();
      //expect(instance).to.be.a(FirstDataGateway.CarRentalExtraCharges);
    });

    it('should have the property chargeItem (base name: "chargeItem")', function() {
      // uncomment below and update the code to test the property chargeItem
      //var instane = new FirstDataGateway.CarRentalExtraCharges();
      //expect(instance).to.be();
    });

  });

}));
