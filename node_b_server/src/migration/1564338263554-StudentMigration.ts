import {MigrationInterface, QueryRunner} from "typeorm";

export class StudentMigration1564338263554 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("CREATE TABLE `student` (`studentId` int NOT NULL AUTO_INCREMENT, `studentName` varchar(255) NOT NULL, `studentClass` varchar(255) NOT NULL, `studentRollNo` int NOT NULL, PRIMARY KEY (`studentId`)) ENGINE=InnoDB");
    }

    public async down(queryRunner: QueryRunner): Promise<any> {
        await queryRunner.query("DROP TABLE `student`");
    }

}
