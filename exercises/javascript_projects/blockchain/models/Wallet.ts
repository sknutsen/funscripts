import * as crypto from 'crypto';
import { RSAKeyPairOptions } from 'crypto';
import { Chain } from "./Chain";
import { Transaction } from './Transaction';

export class Wallet {
    public publicKey: string;
    public privateKey: string;

    constructor() {
        const type: "rsa" = 'rsa';
        const options: RSAKeyPairOptions<'pem', 'pem'> = {
            modulusLength: 2048,
            publicKeyEncoding: {type: 'spki', format: 'pem'},
            privateKeyEncoding: {type: 'pkcs8', format: 'pem'},
        };

        const keypair = crypto.generateKeyPairSync(type, options);

        this.publicKey = keypair.publicKey;
        this.privateKey = keypair.privateKey;
    }

    sendMoney(amount: number, payeePublicKey: string) {
        const transaction = new Transaction(amount, this.publicKey, payeePublicKey);

        const sign = crypto.createSign('SHA256');
        sign.update(transaction.toString()).end();

        const signature = sign.sign(this.privateKey);
        Chain.instance.addBlock(transaction, this.publicKey, signature);
    }
}