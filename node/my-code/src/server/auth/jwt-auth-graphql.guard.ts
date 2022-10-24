import { ExecutionContext, Injectable } from '@nestjs/common';
import { GqlExecutionContext } from '@nestjs/graphql';
import { AuthGuard } from '@nestjs/passport';

@Injectable()
export class JwtAuthGraphqlGuard extends AuthGuard('jwt') {
    getRequest(context: ExecutionContext) {
        const ctx = GqlExecutionContext.create(context);
        const req = ctx.getContext().req;
        req.body = context.switchToHttp().getResponse();
        return req;
    }
}
