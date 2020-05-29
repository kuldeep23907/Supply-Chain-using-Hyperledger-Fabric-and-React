const network = require('../fabric/network.js');
const apiResponse = require('../utils/apiResponse.js');

exports.signup = async (isManufacturer, isMiddlemen, isConsumer, information) => {
    const { userType, address, name, email } = information;

    let networkObj;
    networkObj = await network.connect(true, false, false, 'admin');
    
    let contractRes;
    contractRes = await network.invoke(networkObj, 'createUser', name, email, userType, address);
    console.log('5');
    const walletRes = await network.registerUser(isManufacturer, isMiddlemen, isConsumer, contractRes.ID);

    const error = walletRes.error || networkObj.error || contractRes.error;
    if (error) {
        const status = walletRes.status || networkObj.status || contractRes.status;
        return apiResponse.createModelRes(status, error);
    }

    return apiResponse.createModelRes(200, 'Success', contractRes);
};