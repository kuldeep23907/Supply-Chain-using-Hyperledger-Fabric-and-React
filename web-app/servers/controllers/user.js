const authModel = require('../models/auth.js');
const apiResponse = require('../utils/apiResponse.js');

exports.signup = async (req, res) => {
    const { userType, address, name, email } = req.body;
    const { role } = req.params;

    console.log(req.body);
    console.log(role);

    if ((!userType || !address || !name  || !email )) {
        console.log('1');
        return apiResponse.badRequest(res);
    }

    let modelRes;

    if (role === 'manufacturer') {
        modelRes = await authModel.signup(true, false, false, {  userType, address, name, email });
    } else if (role === 'middlemen') {
        modelRes = await authModel.signup(false, true, false, {  userType, address, name, email });
    } else if (role === 'consumer') {
        modelRes = await authModel.signup(false, false, true, {  userType, address, name, email });
    } else {
        return apiResponse.badRequest(res);
    }

    return apiResponse.send(res, modelRes);
};