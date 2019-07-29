import { Router } from "express";
import { StudentApi } from "./student.api";

const router: Router = Router();

router.use('/student', StudentApi);

export const v1Api: Router = router;
