import Sequelize from 'sequelize';


let sequelize;
const env = process.env.NODE_ENV || 'development';

if (process.env.SEQUELIZE_URI) {
    sequelize = new Sequelize(process.env.SEQUELIZE_URI, {
        paranoid: true,
    });
}
