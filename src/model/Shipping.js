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
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Address', 'model/Contact'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Address'), require('./Contact'));
  } else {
    // Browser globals (root is window)
    if (!root.FirstDataGateway) {
      root.FirstDataGateway = {};
    }
    root.FirstDataGateway.Shipping = factory(root.FirstDataGateway.ApiClient, root.FirstDataGateway.Address, root.FirstDataGateway.Contact);
  }
}(this, function(ApiClient, Address, Contact) {
  'use strict';




  /**
   * The Shipping model module.
   * @module model/Shipping
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Shipping</code>.
   * @alias module:model/Shipping
   * @class
   */
  var exports = function() {
    var _this = this;




  };

  /**
   * Constructs a <code>Shipping</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Shipping} obj Optional instance to populate.
   * @return {module:model/Shipping} The populated <code>Shipping</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
      if (data.hasOwnProperty('contact')) {
        obj['contact'] = Contact.constructFromObject(data['contact']);
      }
      if (data.hasOwnProperty('address')) {
        obj['address'] = Address.constructFromObject(data['address']);
      }
    }
    return obj;
  }

  /**
   * @member {String} name
   */
  exports.prototype['name'] = undefined;
  /**
   * @member {module:model/Contact} contact
   */
  exports.prototype['contact'] = undefined;
  /**
   * @member {module:model/Address} address
   */
  exports.prototype['address'] = undefined;



  return exports;
}));


