import { Chain } from './models/Chain';
import { Wallet } from './models/Wallet';

const satoshi = new Wallet();
const human = new Wallet();
const alien = new Wallet();

satoshi.sendMoney(50, human.publicKey);
human.sendMoney(25, alien.publicKey);
alien.sendMoney(5, human.publicKey);

console.log(Chain.instance);