import bugsnag from 'bugsnag';
import grpc from 'grpc';
import loader from '@grpc/proto-loader';

bugsnag.register(process.env.BUGSNAG_API)

const server = new grpc.Server();
