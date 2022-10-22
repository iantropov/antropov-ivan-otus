import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ViewModule } from './view/view.module';
import { UsersModule } from './users/users.module';

@Module({
    imports: [ViewModule, MongooseModule.forRoot('mongodb://localhost:27017/my-code'), UsersModule],
    controllers: [AppController],
    providers: [AppService]
})
export class AppModule {}
