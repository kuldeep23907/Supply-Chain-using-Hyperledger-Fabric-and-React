const productRouter = require('express').Router();
const controller = require('../controllers/product.js');

productRouter.post('/', controller.createProduct);
productRouter.put('/:productId', controller.updateProduct);
productRouter.get('/:productId', controller.getProductbyId);
productRouter.get('/', controller.getAllProducts);

const authMiddleware = require('../middlewares/auth.js');

productRouter.use('/order', authMiddleware);

productRouter.post('/order', controller.createOrder);


module.exports = entityRouter;