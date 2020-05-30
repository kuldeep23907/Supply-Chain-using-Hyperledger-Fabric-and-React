const apiResponse = require('../utils/apiResponse.js');

module.exports = async (req, res, next) => {
    const {userType} = req.body;
    console.log(req.body);

    if (!userType) {
        return apiResponse.unauthorized(res, 'Unauthorised user');
    }

    try {
        if( userType === 'admin') {
            return next();
        }
        return apiResponse.unauthorized(res, "User type admin required");
    } catch (err) {
        return apiResponse.unauthorized(res, err.toString());
    }
};
