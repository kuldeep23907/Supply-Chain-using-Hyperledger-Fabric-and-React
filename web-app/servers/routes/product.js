const productRouter = require('express').Router();
const controller = require('../controllers/product.js');
const authMiddleware = require('../middlewares/auth.js');
const roleMiddleware = require('../middlewares/checkRole.js');


productRouter.use('/', authMiddleware);

productRouter.post('/', controller.createProduct);
productRouter.put('/:productId/:role', controller.updateProduct);
productRouter.get('/:productId/:role', controller.getProductbyId);
productRouter.get('/:role', controller.getAllProducts);

module.exports = productRouter;