export class Transaction {
    constructor(public ammount: number, public payer: string, public payee: string) {}

    public toString() {
        return JSON.stringify(this);
    }
}