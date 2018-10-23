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
    define(['ApiClient', 'model/AirlineAncillaryServiceCategory', 'model/AirlineTravelRoute'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./AirlineAncillaryServiceCategory'), require('./AirlineTravelRoute'));
  } else {
    // Browser globals (root is window)
    if (!root.FirstDataGateway) {
      root.FirstDataGateway = {};
    }
    root.FirstDataGateway.Airline = factory(root.FirstDataGateway.ApiClient, root.FirstDataGateway.AirlineAncillaryServiceCategory, root.FirstDataGateway.AirlineTravelRoute);
  }
}(this, function(ApiClient, AirlineAncillaryServiceCategory, AirlineTravelRoute) {
  'use strict';




  /**
   * The Airline model module.
   * @module model/Airline
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Airline</code>.
   * Additional data specific to the airline industry
   * @alias module:model/Airline
   * @class
   */
  var exports = function() {
    var _this = this;















  };

  /**
   * Constructs a <code>Airline</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Airline} obj Optional instance to populate.
   * @return {module:model/Airline} The populated <code>Airline</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('passengerName')) {
        obj['passengerName'] = ApiClient.convertToType(data['passengerName'], 'String');
      }
      if (data.hasOwnProperty('ticketNumber')) {
        obj['ticketNumber'] = ApiClient.convertToType(data['ticketNumber'], 'String');
      }
      if (data.hasOwnProperty('issuingCarrier')) {
        obj['issuingCarrier'] = ApiClient.convertToType(data['issuingCarrier'], 'String');
      }
      if (data.hasOwnProperty('carrierName')) {
        obj['carrierName'] = ApiClient.convertToType(data['carrierName'], 'String');
      }
      if (data.hasOwnProperty('travelAgencyIataCode')) {
        obj['travelAgencyIataCode'] = ApiClient.convertToType(data['travelAgencyIataCode'], 'String');
      }
      if (data.hasOwnProperty('travelAgencyName')) {
        obj['travelAgencyName'] = ApiClient.convertToType(data['travelAgencyName'], 'String');
      }
      if (data.hasOwnProperty('airlinePlanNumber')) {
        obj['airlinePlanNumber'] = ApiClient.convertToType(data['airlinePlanNumber'], 'String');
      }
      if (data.hasOwnProperty('airlineInvoiceNumber')) {
        obj['airlineInvoiceNumber'] = ApiClient.convertToType(data['airlineInvoiceNumber'], 'String');
      }
      if (data.hasOwnProperty('reservationSystem')) {
        obj['reservationSystem'] = ApiClient.convertToType(data['reservationSystem'], 'String');
      }
      if (data.hasOwnProperty('restricted')) {
        obj['restricted'] = ApiClient.convertToType(data['restricted'], 'Boolean');
      }
      if (data.hasOwnProperty('travelRoute')) {
        obj['travelRoute'] = ApiClient.convertToType(data['travelRoute'], [AirlineTravelRoute]);
      }
      if (data.hasOwnProperty('relatedTicketNumber')) {
        obj['relatedTicketNumber'] = ApiClient.convertToType(data['relatedTicketNumber'], 'String');
      }
      if (data.hasOwnProperty('ancillaryServiceCategory')) {
        obj['ancillaryServiceCategory'] = ApiClient.convertToType(data['ancillaryServiceCategory'], [AirlineAncillaryServiceCategory]);
      }
      if (data.hasOwnProperty('ticketPurchase')) {
        obj['ticketPurchase'] = ApiClient.convertToType(data['ticketPurchase'], 'Boolean');
      }
    }
    return obj;
  }

  /**
   * @member {String} passengerName
   */
  exports.prototype['passengerName'] = undefined;
  /**
   * @member {String} ticketNumber
   */
  exports.prototype['ticketNumber'] = undefined;
  /**
   * @member {String} issuingCarrier
   */
  exports.prototype['issuingCarrier'] = undefined;
  /**
   * @member {String} carrierName
   */
  exports.prototype['carrierName'] = undefined;
  /**
   * @member {String} travelAgencyIataCode
   */
  exports.prototype['travelAgencyIataCode'] = undefined;
  /**
   * @member {String} travelAgencyName
   */
  exports.prototype['travelAgencyName'] = undefined;
  /**
   * @member {String} airlinePlanNumber
   */
  exports.prototype['airlinePlanNumber'] = undefined;
  /**
   * @member {String} airlineInvoiceNumber
   */
  exports.prototype['airlineInvoiceNumber'] = undefined;
  /**
   * @member {module:model/Airline.ReservationSystemEnum} reservationSystem
   */
  exports.prototype['reservationSystem'] = undefined;
  /**
   * @member {Boolean} restricted
   */
  exports.prototype['restricted'] = undefined;
  /**
   * @member {Array.<module:model/AirlineTravelRoute>} travelRoute
   */
  exports.prototype['travelRoute'] = undefined;
  /**
   * @member {String} relatedTicketNumber
   */
  exports.prototype['relatedTicketNumber'] = undefined;
  /**
   * Identify the purchase of ancillary goods or services with a false value. If this element is not provided, the transaction is assumed to be a purchase of an airline ticket.
   * @member {Array.<module:model/AirlineAncillaryServiceCategory>} ancillaryServiceCategory
   */
  exports.prototype['ancillaryServiceCategory'] = undefined;
  /**
   * @member {Boolean} ticketPurchase
   */
  exports.prototype['ticketPurchase'] = undefined;


  /**
   * Allowed values for the <code>reservationSystem</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ReservationSystemEnum = {
    /**
     * value: "START"
     * @const
     */
    "START": "START",
    /**
     * value: "TWA"
     * @const
     */
    "TWA": "TWA",
    /**
     * value: "DELTA"
     * @const
     */
    "DELTA": "DELTA",
    /**
     * value: "SABRE"
     * @const
     */
    "SABRE": "SABRE",
    /**
     * value: "COVIA_APOLLO"
     * @const
     */
    "COVIA_APOLLO": "COVIA_APOLLO",
    /**
     * value: "DR_BLANK"
     * @const
     */
    "DR_BLANK": "DR_BLANK",
    /**
     * value: "DER"
     * @const
     */
    "DER": "DER",
    /**
     * value: "TUI"
     * @const
     */
    "TUI": "TUI"  };


  return exports;
}));


