'use strict';

const { Contract } = require('fabric-contract-api');

class CarChaincode extends Contract {
    async initLedger(ctx) {
    
        const cars = [
            {
                make: 'Honda',
                carId: '',
                model: 'Civic',
                owner: '',
                state: '',
                currstatus:'',
            },
            {
                make: 'Toyota',
                carId: '',
                model: 'Camry',
                owner: '',
                state: '',
                currstatus: '',
            },
            {
                make: 'Ford',
                carId: '',
                model: 'F150',
                owner: '',
                state: '',
                currstatus: '',
            },
        ];

        for (let i = 0; i < cars.length; i++) {
            cars[i].docType = 'car';
            await ctx.stub.putState('CAR' + i, Buffer.from(JSON.stringify(cars[i])));
            console.info('Added', cars[i]);
        }
        
    }


    async ManufactureCar(ctx, carId, make, model, owner) {

        const carAsBytes = await ctx.stub.getState(carNumber); // get the car from chaincode state
        if (!carAsBytes || carAsBytes.length === 0) {
            throw new Error(`${carNumber} does not exist`);
        }

        if (carId.length < 5){
            throw new Error('invalid id')
        }
        if(owner == 'Manufacturer'){
        const car = JSON.parse(carAsBytes.toString());
        car.make = make;
        car.owner = 'Manufacturer';
        car.model = model;
        car.state = 'CREATED';
        car.currstatus = 'factory';
        await ctx.stub.putState(carNumber, Buffer.from(JSON.stringify(car)));
    }
        else{
            console.log('error owner not a manufacturer');
        }
    }

    async carTransportToDealer(ctx, carId, Drname) {

        const carAsBytes = await ctx.stub.getState(carId); // get the car from chaincode state
        if (!carAsBytes || carAsBytes.length === 0) {
            throw new Error(`${carId} does not exist`);
        }
        const car = JSON.parse(carAsBytes.toString());
        if(car.currstatus == 'factory'){
            car.currstatus = Drname;
            console.log("car is being delivered by ",Drname);
        }
    }

    async carDeliveredToDealer(ctx , carId, name){
        const carAsBytes = await ctx.stub.getState(carId); // get the car from chaincode state
        if (!carAsBytes || carAsBytes.length === 0) {
            throw new Error(`${carId} does not exist`);
        }
        const car = JSON.parse(carAsBytes.toString());
        if (car.currstatus == name)
        {
            car.currstatus = 'received by dealer';
            console.log('car received by dealer');
        }

    }

    async Dealer(ctx, carId, make, model) {

        const carAsBytes = await ctx.stub.getState(carId); // get the car from chaincode state
        if (!carAsBytes || carAsBytes.length === 0) {
            throw new Error(`${carNumber} does not exist`);
        }
        console.log('car is delivered to dealer');
        const car = JSON.parse(carAsBytes.toString());
        car.make = make;
        car.owner = 'Dealer';
        car.model = model;
        car.state = 'READY_FOR_SALE';
        car.currstatus = 'with Dealer';

        await ctx.stub.putState(carNumber, Buffer.from(JSON.stringify(car)));
        
    }

    async sellCar(ctx, carId, customer) {

        const carAsBytes = await ctx.stub.getState(carId); // get the car from chaincode state
        if (!carAsBytes || carAsBytes.length === 0) {
            throw new Error(`${carNumber} does not exist`);
        }
        if (carAsBytes.owner == Dealer){
        const car = JSON.parse(carAsBytes.toString());
        car.owner = customer;
        car.state = 'SOLD';
        car.currstatus = 'Sold';
        await ctx.stub.putState(carNumber, Buffer.from(JSON.stringify(car)));
    }
        else{
            console.log('car is not with dealer');
        }

        
    }

    async queryCar(ctx, carNumber) {
        const carAsBytes = await ctx.stub.getState(carNumber); // get the car from chaincode state
        if (!carAsBytes || carAsBytes.length === 0) {
            throw new Error(`${carNumber} does not exist`);
        }
        console.log(carAsBytes.toString());
        return carAsBytes.toString();
    }
}

module.exports = CarChaincode;
