const router = require('express').Router();

const userRouter = require('./user.js');
const productRouter = require('./product.js');


router.use('/user', userRouter);
router.use('/product', productRouter);

module.exports = router;
