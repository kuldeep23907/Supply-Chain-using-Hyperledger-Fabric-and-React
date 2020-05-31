const apiResponse = require('../utils/apiResponse.js');

module.exports = async (req, res, next) => {
    const {loggedUserType} = req.body;
    console.log(req.body);

    if (!loggedUserType) {
        return apiResponse.unauthorized(res, 'Unauthorised user');
    }

    try {
        if( loggedUserType === 'admin') {
            return next();
        }
        return apiResponse.unauthorized(res, "User type admin required");
    } catch (err) {
        return apiResponse.unauthorized(res, err.toString());
    }
};
