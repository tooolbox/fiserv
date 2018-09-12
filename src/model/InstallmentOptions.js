/**
 * Payment Gateway API Specification
 * Payment Gateway API for payment processing. 
 *
 * OpenAPI spec version: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from '../ApiClient';





/**
* The InstallmentOptions model module.
* @module model/InstallmentOptions
* @version 6.3.0
*/
export default class InstallmentOptions {
    /**
    * Constructs a new <code>InstallmentOptions</code>.
    * Indicates that the total sum payable is divided for payment at successive fixed times.
    * @alias module:model/InstallmentOptions
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>InstallmentOptions</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/InstallmentOptions} obj Optional instance to populate.
    * @return {module:model/InstallmentOptions} The populated <code>InstallmentOptions</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstallmentOptions();

            
            
            

            if (data.hasOwnProperty('numberOfInstallments')) {
                obj['numberOfInstallments'] = ApiClient.convertToType(data['numberOfInstallments'], 'Number');
            }
            if (data.hasOwnProperty('installmentsInterest')) {
                obj['installmentsInterest'] = ApiClient.convertToType(data['installmentsInterest'], 'Boolean');
            }
            if (data.hasOwnProperty('installmentDelayMonths')) {
                obj['installmentDelayMonths'] = ApiClient.convertToType(data['installmentDelayMonths'], 'Number');
            }
        }
        return obj;
    }

    /**
    * Number of installments for a Sale transaction if the customer pays the total amount in multiple transactions
    * @member {Number} numberOfInstallments
    */
    numberOfInstallments = undefined;
    /**
    * Indicates whether the installment interest amount has been applied. Possible values are \"yes\" or \"no\".
    * @member {Boolean} installmentsInterest
    */
    installmentsInterest = undefined;
    /**
    * The number of months the first installment payment will be delayed
    * @member {Number} installmentDelayMonths
    */
    installmentDelayMonths = undefined;








}

