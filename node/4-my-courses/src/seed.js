import mongoose from './mongoose.js';
import { Lesson } from './models/lesson.js';
import { Course } from './models/course.js';

mongoose.start().then(async () => {
    console.log("Removing all previous data...");
    await Course.deleteMany();

    console.log("Seeding data...");
    const lesson1 = new Lesson({
        name: 'lesson1',
        description: 'lesson1 description'
    });
    const lesson2 = new Lesson({
        name: 'lesson2',
        description: 'lesson2 description'
    });
    const course1 = new Course({
        name: 'course1',
        lessons: [lesson1, lesson2]
    })
    await course1.save();

    const lesson3 = new Lesson({
        name: 'lesson3',
        description: 'lesson3 description'
    });
    const lesson4 = new Lesson({
        name: 'lesson4',
        description: 'lesson4 description'
    });
    const course2 = new Course({
        name: 'course2',
        lessons: [lesson3, lesson4]
    })
    await course2.save();

    console.log("Finished seeding.");
    process.exit(0);
});