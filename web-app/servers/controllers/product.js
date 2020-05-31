const entityModel = require('../models/product.js');
const apiResponse = require('../utils/apiResponse.js');

exports.createProduct = async (req, res) => {
    const { name, manufacturerId, price , userType } = req.body;
    console.log('1');

    if (!name || !manufacturerId || !price || !usertype) {
        return apiResponse.badRequest(res);
    }
    console.log('2');

    if (!userType == 'manufacturer' ) {
        return apiResponse.badRequest(res);
    }
    console.log('3');

    const modelRes = await productModel.createProduct({ name, manufacturerId, price });
    return apiResponse.send(res, modelRes);
};

exports.updateProduct = async (req, res) => {
    const { productID, userId, name , price } = req.body;
    const { userType } = req.params;

    if (!productID || !userId || !name || !price) {
        return apiResponse.badRequest(res);
    }

    if (userType == 'consumer') {
        return apiResponse.badRequest(res);
    }

    const modelRes = await entityModel.updateProduct({ productID, userId, name , price });
    return apiResponse.send(res, modelRes);
};

exports.createOrder = async (req, res) => {
    const { productID, userId , userType , name } = req.body;

    if (!productID || !userId || !userType || !name) {
        return apiResponse.badRequest(res);
    }

    if (userType != 'consumer') {
        return apiResponse.badRequest(res);
    }

    const modelRes = await entityModel.createOrder({ productID, userId, userType , name });
    return apiResponse.send(res, modelRes);

};

exports.isDelivered = async (req, res) => {
    // find who initiates this event by decoding the token and getting the user type
    const { id, userType, name , productId , userId } = req.body;
    console.log('1');
    if (!name || !userId || !userType || !productId || !id) {
        return apiResponse.badRequest(res);
    }
    console.log('2');
    const modelRes;
    if(userType == 'consumer')
    {
        modelRes= await transactModel.isDelivered({ productId , id });
    }

    console.log('3');
    return apiResponse.send(res, modelRes);
};
