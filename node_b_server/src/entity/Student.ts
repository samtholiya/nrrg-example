import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";

@Entity()
export class Student {
    @PrimaryGeneratedColumn({
        name:"studentId"
    })
    id: number;

    @Column({
        name:"studentName",
        nullable: false
    })
    name: string;

    @Column({
        name:"studentClass",
        nullable: false
    })
    class: string;

    @Column({
        name: "studentRollNo",
        nullable: false
    })
    rollNo: number;
}