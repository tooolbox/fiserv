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
    define(['ApiClient', 'model/TransactionResponseAuthenticationRedirectParams'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./TransactionResponseAuthenticationRedirectParams'));
  } else {
    // Browser globals (root is window)
    if (!root.FirstDataGateway) {
      root.FirstDataGateway = {};
    }
    root.FirstDataGateway.TransactionResponseAuthenticationRedirect = factory(root.FirstDataGateway.ApiClient, root.FirstDataGateway.TransactionResponseAuthenticationRedirectParams);
  }
}(this, function(ApiClient, TransactionResponseAuthenticationRedirectParams) {
  'use strict';




  /**
   * The TransactionResponseAuthenticationRedirect model module.
   * @module model/TransactionResponseAuthenticationRedirect
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>TransactionResponseAuthenticationRedirect</code>.
   * @alias module:model/TransactionResponseAuthenticationRedirect
   * @class
   */
  var exports = function() {
    var _this = this;




  };

  /**
   * Constructs a <code>TransactionResponseAuthenticationRedirect</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TransactionResponseAuthenticationRedirect} obj Optional instance to populate.
   * @return {module:model/TransactionResponseAuthenticationRedirect} The populated <code>TransactionResponseAuthenticationRedirect</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('type')) {
        obj['type'] = ApiClient.convertToType(data['type'], 'String');
      }
      if (data.hasOwnProperty('targetUrl')) {
        obj['targetUrl'] = ApiClient.convertToType(data['targetUrl'], 'String');
      }
      if (data.hasOwnProperty('params')) {
        obj['params'] = TransactionResponseAuthenticationRedirectParams.constructFromObject(data['params']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/TransactionResponseAuthenticationRedirect.TypeEnum} type
   */
  exports.prototype['type'] = undefined;
  /**
   * @member {String} targetUrl
   */
  exports.prototype['targetUrl'] = undefined;
  /**
   * @member {module:model/TransactionResponseAuthenticationRedirectParams} params
   */
  exports.prototype['params'] = undefined;


  /**
   * Allowed values for the <code>type</code> property.
   * @enum {String}
   * @readonly
   */
  exports.TypeEnum = {
    /**
     * value: "SECURE_3D"
     * @const
     */
    "3D": "SECURE_3D"  };


  return exports;
}));


